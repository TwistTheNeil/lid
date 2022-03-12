<template>
  <div class="home">
    <SearchbarInput />
    <FileList />
  </div>
</template>

<script>
import SearchbarInput from "@/components/SearchbarInput.vue";
import FileList from "@/components/FileList.vue";

export default {
  name: "Home",
  components: {
    SearchbarInput,
    FileList,
  },
};
</script>
