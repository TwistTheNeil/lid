<template>
  <table class="table table-responsive table-borderless">
    <thead class="bg-light">
      <tr>
        <th scope="col" width="10%">Hash</th>
        <th scope="col">Name</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="file in filteredFiles" :key="file.hash">
        <td class="monospace">{{ file.hash }}</td>
        <td>{{ file.name }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import { storeToRefs } from "pinia";
import { computed } from "vue";
import { Fzf } from "fzf";

import { useFileStore } from "../store/fileStore";
import { useSearchStore } from "../store/searchStore";

export default {
  name: "FileList",
  setup() {
    const fileStore = useFileStore();
    const searchStore = useSearchStore();
    const { files } = storeToRefs(fileStore);

    const filteredFiles = computed(() => {
      if (!files || files.value.length === 0) {
        return [];
      }

      if (searchStore.search === "") {
        return files.value;
      }

      const fzf = new Fzf(files.value, {
        selector: (i) => i.name,
      });
      return fzf.find(searchStore.search).map((i) => i.item);
    });

    return {
      filteredFiles,
    };
  },
};
</script>
