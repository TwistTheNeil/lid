<template>
  <div class="form-floating mt-4 mb-4">
    <input
      type="text"
      class="form-control border-0"
      id="searchbar"
      placeholder="Search..."
      v-model="search"
    />
    <label for="searchBar">Search...</label>
  </div>
</template>

<script>
import { ref, watch } from "vue";

import { useSearchStore } from "../store/searchStore";

export default {
  setup() {
    let search = ref("");
    const searchStore = useSearchStore();

    watch(
      () => search.value,
      // eslint-disable-next-line no-unused-vars
      (search, oldSearch) => {
        searchStore.search = search;
      }
    );

    return {
      search,
    };
  },
};
</script>
