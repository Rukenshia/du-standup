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
          Starting Time <input type="datetime-local" placeholder="Start Time" v-model="start" />
          <input type="text" placeholder="Where?" v-model="where" />
        </div>
        <div class="column">
          <button class="button-black" @click="add">Add</button>
        </div>
      </div>
    </div>`,
  computed: {
    standupDate() {
      return moment(this.$store.state.standup.Expires);
    }
  },
  mounted() {
    this.start = this.standupDate.format("YYYY-MM-DDTHH:mm");
  },
  methods: {
    add() {
      if (this.title.length === 0 || this.start.length === 0 || this.where.length === 0) {
        this.invalid = true;
        return;
      }
      this.invalid = false;

      this.$emit('add', {
        Title: this.title,
        Start: this.start,
        Where: this.where,
      })
    }
  }
});
