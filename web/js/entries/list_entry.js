Vue.component('list-entry', {
  props: ['title', 'votes'],
  data() {
    return {};
  },
  template: `
        <li>{{ title }}
          (+{{votes}}
          <button class="button-small button-black" @click="vote">+</button>)
          <button class="button-small button-black" @click="deleteEntry">x</button>
        </li>`,
  methods: {
    vote() {
      this.$emit('vote');
    },
    deleteEntry() {
      this.$emit('delete');
    },
  },
})