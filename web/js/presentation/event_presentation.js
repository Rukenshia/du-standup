Vue.component('event-presentation', {
  props: ['events'],
  data() {
    return {
    };
  },
  template: `
    <div class="category-entries">
      <div v-for="event in today" class="row">
        <div class="column column-25 column-offset-10">
          <strong>{{ moment(event.Start).format('HH:mm') }}</strong>&nbsp;
        </div>
        <div class="column event-title" :class="{ 'small': event.Title.length > 30, 'xsmall': event.Title.length > 45 }">
          {{ event.Title }}
        </div>
        <div class="column">
          <small>{{ event.Where }}</small>
        </div>
      </div>
      <hr v-if="later.length > 0 && today.length > 0" />
      <div v-for="event in later" class="row">
        <div class="column column-10">
          <small>{{ moment(event.Start).format('DD.MM') }}</small>
        </div>
        <div class="column column-25">
          <strong>{{ moment(event.Start).format('HH:mm') }}</strong>
        </div>
        <div class="column event-title" :class="{ 'small': event.Title.length > 30, 'xsmall': event.Title.length > 45 }">
          {{ event.Title }}
        </div>
        <div class="column">
          <small>{{ event.Where }}</small>
        </div>
      </div>
    </div>`,
  computed: {
    standupDate() {
      return this.$store.state.standup.Expires;
    },
    moment() {
      return moment;
    },
    today() {
      const date = moment(this.standupDate).date();
      console.log(date);
      return this.events.filter(ev => {
        console.log(moment(ev.Start).date());
        console.log(ev);
        return moment(ev.Start).date() === date;
      });
    },
    later() {
      const date = moment(this.standupDate).date();
      return this.events.filter(ev => {
        return moment(ev.Start).date() > date;
      });
    },
  },

});
