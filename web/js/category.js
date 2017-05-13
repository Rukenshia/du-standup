Vue.component('category', {
  props: ['id', 'type', 'name', 'entries'],
  data() {
    return {
      title: '',
      where: '',
      start: '',
    };
  },
  template: `
  <div class="category">
    <h1>{{ name }}</h1>

    <ul v-if="type === 'list'">
        <li v-for="entry in entries">{{ entry.Title }}
          (+{{entry.Votes}}
          <button class="button-small button-black" @click="voteEntry(entry)">+</button>)
          <button class="button-small button-black" @click="deleteEntry(entry)">x</button>
        </li>
    </ul>
    <div v-if="type === 'events'">
      <div v-for="event in entries">
        {{ event.Title }} starting at {{ event.Start }} in {{ event.Where }}
          <button class="button-small button-black" @click="voteEntry(event)">vote</button>
          <button class="button-small button-black" @click="deleteEntry(event)">x</button>
      </div>
    </div>

    <div class="row">
      <div class="column column-40">
        <input type="text" placeholder="Title" v-model="title" />

        <div v-if="type === 'events'">
          Starting Time <input type="time" placeholder="Start Time" v-model="start" />
          <input type="text" placeholder="Where?" v-model="where" />
        </div>
      </div>
      <div class="column">
        <button class="button-black" @click="addEntry()">Add</button>
      </div>
    </div>
  </div>`,

  methods: {
    generateUrl(endpoint, entryId = null) {
      return `${window.baseURL}/api/categories/${this.id}/${endpoint}${entryId !== null ? '/' + entryId : ''}`;
    },
    addEntry() {
      if (this.title.length === 0) {
        return;
      }

      const postBody = {
        Title: this.title,
      };

      if (this.type === 'events') {
        if (this.start.length === 0 || this.where.length === 0) {
          return;
        }

        // set start time
        const timeSplit = this.start.split(':').map(x => parseInt(x, 10));
        const start = moment(this.$store.state.standup.Expires).hours(timeSplit[0]).minutes(timeSplit[1]).toDate();


        if (this.type === 'events') {
          postBody.Start = start;
          postBody.Where = this.where;
        }
      }

      http.post(this.generateUrl('entries'), JSON.stringify(postBody)).then(body => {
        this.$store.commit('add_entry', { categoryId: this.id, entry: JSON.parse(body) });
      });
    },
    voteEntry(entry) {
      if (localStorage.getItem(`${getDate()}_${entry.ID}_${this.Name}_${entry.Title}`, 'voted')) {
        return;
      }

      localStorage.setItem(`${getDate()}_${entry.ID}_${this.Name}_${entry.Title}`, 'voted');
      http.post(`${this.generateUrl('entries', entry.ID)}/vote`, JSON.stringify({
        Category: entry.Category,
        Title: entry.Title,
        Votes: entry.Votes + 1,
      })).then(body => {
        const entry = JSON.parse(body);
        this.$store.commit('add_entry', { categoryId: this.id, entry });
      });
    },
    deleteEntry(entry) {
      http.del(this.generateUrl('entries', entry.ID))
        .then(body => {
          this.$store.commit('delete_entry', { categoryId: this.id, entry });
        });
    },
  }
});
