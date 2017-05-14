Vue.component('event-entry', {
  props: ['title', 'start', 'location'],
  data() {
    return {};
  },
  template: `
        <div>
          {{ title }} starting at {{ start }} in {{ location }}
          <button class="button-small button-black" @click="deleteEntry">x</button>
        </div>`,
  methods: {
    deleteEntry() {
      this.$emit('delete');
    },
  },
})