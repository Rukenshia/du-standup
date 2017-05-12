var app = new Vue({
  el: '#app',
  store: window.store,
  data: {
    presentationMode: true,
    selectedCategory: null,
  },
  mounted() {
    this.getStandup();
  },
  computed: {
    standup() {
      return this.$store.state.standup;
    },
    categories() {
      return this.$store.state.standup.Categories;
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
          console.log(body);
          this.$store.commit('set_standup', body);
        });
    },
    viewCategory(c) {
      this.selectedCategory = c;
    }
  },
});

function getDate() {
  const d = new Date();

  return `${d.getFullYear()}-${d.getMonth()}-${d.getDay()}`;
}