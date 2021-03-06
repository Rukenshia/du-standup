var app = new Vue({
  el: '#app',
  store: window.store,
  data: {
    presentationMode: true,
    selectedCategory: null,

    date: '0000-00-00',
  },
  mounted() {
    this.getStandup();

    setInterval(() => {
      if (this.presentationMode) {
        // only update current category to reduce load
        this.updateCategoryEntries(this.selectedCategory);
        return;
      }
      this.categories.forEach(c => {
        this.updateCategoryEntries(c);
      });
    }, 3000);

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
  watch: {
    standup(s) {
      this.date = moment(s.Expires).format('YYYY-MM-DD');
    },
  },
  methods: {
    togglePresentationMode() {
      this.presentationMode = !this.presentationMode;

      if (this.presentationMode) {
        this.selectedCategory = this.categories[0];
      }
    },
    updateCategoryEntries(category) {
      http.get(`${window.baseURL}/api/categories/${category.ID}/entries`)
        .then(body => {
          const entries = JSON.parse(body);

          category.Entries = entries;

          if (category.Type === 'list') {
            category.Entries.sort((a, b) => b.Votes - a.Votes);
          } else if (category.Type === 'events') {
            category.Entries.sort((a, b) => moment(a.Start).diff(b.Start));
          }
        });
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
      this.updateCategoryEntries(c);
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