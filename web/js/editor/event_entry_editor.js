Vue.component('event-entry-editor', {
  props: ['id', 'type', 'name', 'events'],
  data() {
    return {
      title: '',
      where: '',
      start: '',
      invalid: false,
    };
  },
  template: `
    <div>
      <div class="row">
        <div class="column">
          <form-error :invalid="invalid"></form-error>
        </div>
      </div>
      <div class="row">
        <div class="column column-40">
          <input type="text" placeholder="Title" v-model="title" />
          Starting Time <input type="time" placeholder="Start Time" v-model="start" />
          <input type="text" placeholder="Where?" v-model="where" />
        </div>
        <div class="column">
          <button class="button-black" @click="add">Add</button>
        </div>
      </div>
    </div>`,

  methods: {
    add() {
      if (this.title.length === 0 || this.start.length === 0 || this.where.length === 0) {
        this.invalid = true;
        return;
      }
      this.invalid = false;

      const timeSplit = this.start.split(':').map(x => parseInt(x, 10));
      const start = moment(this.$store.state.standup.Expires).hours(timeSplit[0]).minutes(timeSplit[1]).toDate();

      this.$emit('add', {
        Title: this.title,
        Start: start,
        Where: this.where,
      })
    }
  }
});
