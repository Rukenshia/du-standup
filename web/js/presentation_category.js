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
        <ul class="category-entries">
            <li v-if="type === 'list'" v-for="entry in entries">
                {{ entry.Title }}
            </li>
        </ul>

        <div v-if="type === 'events'" class="category-entries">
          <div v-for="event in entries" class="row">
            <div class="column column-25">
              <strong>{{ moment(event.Start).format('HH:mm') }}</strong>&nbsp;
            </div>
            <div class="column">
              {{ event.Title }}
            </div>
            <div class="column">
              <small>{{ event.Where }}</small>
            </div>
          </div>
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
