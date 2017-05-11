Vue.component('category', {
  props: ['name', 'entries', 'allowUpdates'],
  data() {
    return {
      input: '',
    };
  },
  template: `
  <div class="category">
    <h1>{{ name }}</h1>
    <ul>
        <li v-for="entry in entries">{{ entry.Title }} (+{{entry.Votes}} <button @click="voteEntry(entry)">+</button>) <button @click="deleteEntry(entry)">x</button></li>
        <li v-if="allowUpdates"><input type="text" v-model="input" /> <button @click="addEntry()">Add</button></li>
    </ul>
  </div>`,

  methods: {
    addEntry() {
      http.post(`${window.baseURL}/api/entries`, JSON.stringify({
        Category: this.name,
        Title: this.input,
      })).then(body => {
        this.$store.commit('add_entry', JSON.parse(body));
      });
    },
    voteEntry(entry) {
      if (localStorage.getItem(`${getDate()}_${entry.ID}`, 'voted')) {
        return;
      }

      localStorage.setItem(`${getDate()}_${entry.ID}`, 'voted');
      http.put(`${window.baseURL}/api/entries/${entry.ID}`, JSON.stringify({
        Category: entry.Category,
        Title: entry.Title,
        Votes: entry.Votes + 1,
      })).then(body => {
        this.$store.commit('add_entry', JSON.parse(body));
      });
    },
    deleteEntry(entry) {
      http.del(`${window.baseURL}/api/entries/${entry.ID}`)
        .then(body => {
          this.$store.commit('delete_entry', entry);
        });
    },
  }
});
