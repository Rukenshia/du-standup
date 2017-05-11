var app = new Vue({
  el: '#app',
  store: window.store,
  data: {
    presentationMode: false,
  },
  mounted() {
    this.getStandup();
  },
  computed: {
    categories() {
      return this.$store.state.categories;
    }
  },
  methods: {
    togglePresentationMode() {
      this.presentationMode = !this.presentationMode;
    },
    getStandup(c) {
      http.get(`${window.baseURL}/api/standup`)
        .then(body => {
          body = JSON.parse(body);
          body.Entries.forEach(e => {
            console.log(e);
            this.$store.commit('add_entry', e);
          });
        });
    },
  },
});

function getDate() {
  const d = new Date();

  return `${d.getFullYear()}-${d.getMonth()}-${d.getDay()}`;
}