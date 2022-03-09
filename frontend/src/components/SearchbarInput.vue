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
export default {
  data() {
    return {
      search: "",
    };
  },
};
</script>
