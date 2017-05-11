const parts = window.location.href.split("/");

const baseURL = `${parts[0]}//${parts[2]}`;


var app = new Vue({
  el: '#app',
  data: {
    standup: {},
    categories: [{
      name: 'needs',
      entries: [],
      input: '',
    }],
  },
  mounted() {
    this.getStandup();
  },
  methods: {
    getStandup(c) {
      http.get(`${baseURL}/api/standup`)
        .then(body => {
          body = JSON.parse(body);
          body.Entries.forEach(this.parseEntry.bind(this));
        });
    },
    parseEntry(e) {
      this.categories.forEach(c => {
        if (c.name === e.Category) {

          const existing = c.entries.find(ee => ee.ID === e.ID);
          if (existing) {
            Object.assign(existing, e);
          } else {
            c.entries.push(e);
          }
        }
      });
    },
    addEntry(category, title) {
      http.post(`${baseURL}/api/entries`, JSON.stringify({
        Category: category,
        Title: title,
      })).then(body => {
        this.parseEntry(JSON.parse(body));
      });
    },
    voteEntry(entry) {
      if (localStorage.getItem(`${getDate()}_${entry.ID}`, 'voted')) {
        return;
      }

      localStorage.setItem(`${getDate()}_${entry.ID}`, 'voted');
      http.put(`${baseURL}/api/entries/${entry.ID}`, JSON.stringify({
        Category: entry.Category,
        Title: entry.Title,
        Votes: entry.Votes + 1,
      })).then(body => {
        this.parseEntry(JSON.parse(body));
      });
    },
    deleteEntry(entry) {
      http.del(`${baseURL}/api/entries/${entry.ID}`)
        .then(body => {
          const cat = this.categories.find(c => entry.Category === c.name);

          if (cat) {
            cat.entries = cat.entries.filter(e => e.ID !== entry.ID);
          }
        });
    },
  },
});

function getDate() {
  const d = new Date();

  return `${d.getFullYear()}-${d.getMonth()}-${d.getDay()}`;
}