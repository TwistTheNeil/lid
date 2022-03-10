<template>
  <table class="table table-striped table-hover table-sm table-bordered mt-5">
    <thead>
      <tr>
        <th v-for="header in headers" scope="col" :key="header">
          {{ header }}
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="file in filteredFiles" :key="file.hash">
        <td>{{ file.hash }}</td>
        <td>{{ file.name }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import { storeToRefs } from "pinia";
import { computed } from "vue";
import Fuse from "fuse.js";

import { useFileStore } from "../store/fileStore";
import { useSearchStore } from "../store/searchStore";

export default {
  name: "FileList",
  props: {
    headers: Array,
  },
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

      const fuseOpts = {
        includeScore: true,
        ignoreLocation: true,
        keys: ["name", "hash"],
      };
      const fuse = new Fuse(files.value, fuseOpts);

      return fuse.search(searchStore.search).map(i => i.item);
    });

    return {
      filteredFiles,
    };
  },
};
</script>
