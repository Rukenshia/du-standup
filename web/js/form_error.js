Vue.component('form-error', {
  props: ['invalid'],
  data() {
    return {};
  },
  template: `
    <blockquote class="error" v-if="invalid">
    Invalid form data.
    </blockquote>`,
});
