<template>
  <div class="home">
    <SearchbarInput />

    <span class="mx-2">
      <small>
        Tracking {{ files.length }} file{{ files.length === 1 ? "" : "s" }}
      </small>
    </span>
    <FileList />
  </div>
</template>

<script>
import { storeToRefs } from "pinia";

import SearchbarInput from "@/components/SearchbarInput.vue";
import FileList from "@/components/FileList.vue";
import { useFileStore } from "@/store/fileStore";

export default {
  name: "Home",
  components: {
    SearchbarInput,
    FileList,
  },
  setup() {
    const fileStore = useFileStore();
    const { files } = storeToRefs(fileStore);
    return {
      files,
    };
  },
};
</script>
