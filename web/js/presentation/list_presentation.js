Vue.component('list-presentation', {
  props: ['entries'],
  data() {
    return {
      input: '',
    };
  },
  template: `
    <ul class="category-entries">
      <li v-for="entry in entries">
        {{ entry.Title }}
      </li>
    </ul>`,

  computed: {
    moment() {
      return moment;
    }
  },

});
