window.store = new Vuex.Store({
  state: {
    categories: [{
      name: 'needs',
      entries: [],
      input: '',
    }],
  },
  mutations: {
    add_entry(state, entry) {
      state.categories.forEach(c => {
        if (c.name === entry.Category) {
          const existing = c.entries.find(ee => ee.ID === entry.ID);
          if (existing) {
            Object.assign(existing, entry);
          } else {
            c.entries.push(entry);
          }
        }
      });
    },
    delete_entry(state, entry) {
      const cat = state.categories.find(c => entry.Category === c.name);

      if (cat) {
        cat.entries = cat.entries.filter(e => e.ID !== entry.ID);
      }
    }
  }
})
