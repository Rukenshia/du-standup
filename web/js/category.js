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
      <template v-for="entry in entries">
        <list-entry :title="entry.Title" :votes="entry.Votes" @vote="voteEntry(entry)" @delete="deleteEntry(entry)"></list-entry>
      </template>

      <list-entry-editor @add="addEntry"></list-entry-editor>
    </ul>
    <div v-if="type === 'events'">
      <div v-for="event in entries">
        <event-entry :title="event.Title" :start="event.Start" :location="event.Where" @delete="deleteEntry(event)"></event-entry>
      </div>

      <event-entry-editor @add="addEntry"></event-entry-editor>
    </div>
  </div>`,

  methods: {
    generateUrl(endpoint, entryId = null) {
      return `${window.baseURL}/api/categories/${this.id}/${endpoint}${entryId !== null ? '/' + entryId : ''}`;
    },
    addEntry(data) {
      http.post(this.generateUrl('entries'), JSON.stringify(data)).then(body => {
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
