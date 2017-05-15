Vue.component('list-entry-editor', {
  data() {
    return {
      title: '',
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
        </div>
        <div class="column">
          <button class="button-black" @click="add">Add</button>
        </div>
      </div>
    </div>`,

  methods: {
    add() {
      if (this.title.length === 0) {
        this.invalid = true;
        return;
      }
      this.invalid = false;

      this.$emit('add', {
        Title: this.title,
      });
    }
  }
});
