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
      <tr v-for="file in files" :key="file.hash">
        <td>{{ file.hash }}</td>
        <td>{{ file.name }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import { storeToRefs } from "pinia";
import { useFileStore } from "../store/fileStore";

export default {
  name: "FileList",
  props: {
    headers: Array,
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
