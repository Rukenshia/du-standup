Vue.component('category', {
  props: ['id', 'name', 'entries'],
  data() {
    return {
      input: '',
    };
  },
  template: `
  <div class="category">
    <h1>{{ name }}</h1>

    <ul>
        <li v-for="entry in entries">{{ entry.Title }}
          (+{{entry.Votes}}
          <button class="button-small button-black" @click="voteEntry(entry)">+</button>)
          <button class="button-small button-black" @click="deleteEntry(entry)">x</button>
        </li>
    </ul>

    <div class="row">
      <div class="column column-40">
        <input type="text" v-model="input" />
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
      if (this.input.length === 0) {
        return;
      }

      http.post(this.generateUrl('entries'), JSON.stringify({
        Category: this.name,
        Title: this.input,
      })).then(body => {
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
