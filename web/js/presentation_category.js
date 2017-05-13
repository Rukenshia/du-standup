Vue.component('presentation-category', {
  props: ['entries'],
  data() {
    return {
      input: '',
    };
  },
  template: `
  <div class="category">
    <div v-if="entries.length === 0" class="category-empty">
      <p>nothing &#x1F389</p>
    </div>
    <ul>
        <li v-for="entry in entries">
            {{ entry.Title }}
        </li>
    </ul>

  </div>`,

});
