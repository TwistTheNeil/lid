<template>
  <table class="table table-responsive table-borderless table-hover">
    <thead class="bg-light">
      <tr>
        <th scope="col" width="10%">Hash</th>
        <th scope="col" width="10%">Size</th>
        <th scope="col" width="10%">Devices</th>
        <th scope="col">Name</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="file in filteredFiles" :key="file.hash">
        <td class="monospace">{{ file.hash }}</td>
        <td class="monospace">{{ file.size }}</td>
        <td class="monospace">
          <span
            v-for="device in file.devices"
            :key="device.uuid"
            class="badge rounded-pill bg-success"
          >
            {{ device.name }}
          </span>
        </td>
        <td>{{ file.name }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import { storeToRefs } from "pinia";
import { computed } from "vue";
import { Fzf } from "fzf";

import { useFileStore } from "@/store/fileStore";
import { useSearchStore } from "@/store/searchStore";
import { bytesToHumanReadableUnits } from "@/services/dataSizeConversion";

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

      const f = files.value.map((f) => ({
        ...f,
        size: bytesToHumanReadableUnits(f.size),
      }));

      if (searchStore.search === "") {
        return f;
      }

      const fzf = new Fzf(f, {
        selector: (i) => i.name,
      });

      // We remove whitespace since fzf won't find a whitespace in certain
      // filenames. As a result, weights won't be accurately calculated
      // and expected results will not be seen
      return fzf.find(searchStore.search.replace(/\s/g, "")).map((i) => i.item);
    });

    return {
      filteredFiles,
    };
  },
};
</script>
