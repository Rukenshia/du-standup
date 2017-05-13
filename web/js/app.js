var app = new Vue({
  el: '#app',
  store: window.store,
  data: {
    presentationMode: true,
    selectedCategory: null,
  },
  mounted() {
    this.getStandup();

    document.addEventListener('keydown', e => {
      if (!this.presentationMode) {
        return;
      }

      if (e.keyCode === 37) {
        // left arrow
        this.previousCategory();
      } else if (e.keyCode === 39) {
        // right arrow
        this.nextCategory();
      }
    });
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

      if (this.presentationMode) {
        this.selectedCategory = this.categories[0];
      }
    },
    getStandup(c) {
      http.get(`${window.baseURL}/api/standup`)
        .then(body => {
          body = JSON.parse(body);
          console.log(body);
          this.$store.commit('set_standup', body);
          this.selectedCategory = this.categories[0];
        });
    },
    viewCategory(c) {
      this.selectedCategory = c;
    },
    nextCategory() {
      const idx = this.categories.indexOf(this.selectedCategory);

      if (idx + 1 >= this.categories.length) {
        this.viewCategory(this.categories[0]);
      } else {
        this.viewCategory(this.categories[idx+1]);
      }
    },
    previousCategory() {
      const idx = this.categories.indexOf(this.selectedCategory);

      if (idx - 1 < 0) {
        this.viewCategory(this.categories[this.categories.length - 1]);
      } else {
        this.viewCategory(this.categories[idx-1]);
      }
    }
  },
});

function getDate() {
  const d = new Date();

  return `${d.getFullYear()}-${d.getMonth()}-${d.getDay()}`;
}