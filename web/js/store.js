window.store = new Vuex.Store({
  state: {
    standup: {},
  },
  mutations: {
    set_standup(state, standup) {
      state.standup = standup;

      standup.Categories.forEach(category => {
        if (category.Type === 'list') {
          category.Entries.sort((a, b) => b.Votes - a.Votes);
        } else if (category.Type === 'events') {
          category.Entries.sort((a, b) => moment(a.Start).diff(b.Start));
        }
      });
    },
    add_category(state, category) {
      if (state.standup.Categories.find(c => c.ID === category.ID)) {
        throw new Error('Category with this ID exists');
      }

      state.standup.Categories.push(category);

      if (category.Type === 'list') {
        category.Entries.sort((a, b) => b.Votes - a.Votes);
      } else if (category.Type === 'events') {
        category.Entries.sort((a, b) => moment(a.Start).diff(b.Start));
      }
    },
    add_entry(state, { categoryId, entry }) {
      const category = state.standup.Categories.find(c => c.ID === categoryId);
      if (!category) {
        console.warn('invalid category, skipping');
        return;
      }

      if (entry.Start) {
        entry.Start = new Date(entry.Start);
      }

      const existing = category.Entries.find(ee => ee.ID === entry.ID);
      if (existing) {
        Object.assign(existing, entry);
      } else {
        category.Entries.push(entry);
      }

      if (category.Type === 'list') {
        category.Entries.sort((a, b) => b.Votes - a.Votes);
      } else if (category.Type === 'events') {
        category.Entries.sort((a, b) => moment(a.Start).diff(b.Start));
      }
    },
    delete_entry(state, { categoryId, entry }) {
      const category = state.standup.Categories.find(c => c.ID === categoryId);
      if (!category) {
        console.warn('invalid category, skipping');
        return;
      }

      category.Entries = category.Entries.filter(e => e.ID !== entry.ID);
    }
  }
})
