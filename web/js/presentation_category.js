Vue.component('presentation-category', {
  props: ['entries', 'type'],
  data() {
    return {
      input: '',
    };
  },
  template: `
  <div class="column category">
    <div class="row">
      <div class="column">
        <div v-if="entries.length === 0" class="category-empty">
          <p>nothing &#x1F389</p>
        </div>

        <template v-if="type === 'list'">
          <list-presentation :entries="entries"></list-presentation>
        </template>
        <template v-else-if="type === 'events'">
          <event-presentation :events="entries"></event-presentation>
        </template>
        </div>
      </div>
    </div>

  </div>`,

  computed: {
    moment() {
      return moment;
    }
  },

});
