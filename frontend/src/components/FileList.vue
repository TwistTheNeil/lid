<template>
  <table class="table table-striped table-hover table-sm">
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
import { useFileStore } from "../store/fileStore";
import { useSearchStore } from "../store/searchStore";
import { computed } from "vue";

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

      return files.value.filter((f) =>
        f.name.toLowerCase().match(searchStore.search.toLowerCase())
      );
    });

    return {
      filteredFiles,
    };
  },
};
</script>
