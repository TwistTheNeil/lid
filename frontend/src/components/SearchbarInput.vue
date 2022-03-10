<template>
  <div class="form-floating">
    <input
      type="text"
      class="form-control"
      id="searchbar"
      placeholder="Search..."
      v-model="search"
    />
    <label for="searchBar">Search...</label>
  </div>
</template>

<script>
import { useSearchStore } from "../store/searchStore";
import { ref, watch } from "vue";

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
