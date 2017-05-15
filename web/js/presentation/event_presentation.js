Vue.component('event-presentation', {
  props: ['events'],
  data() {
    return {
      input: '',
    };
  },
  template: `
    <div class="category-entries">
      <div v-for="event in events" class="row">
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
    </div>`,

  computed: {
    moment() {
      return moment;
    }
  },

});
