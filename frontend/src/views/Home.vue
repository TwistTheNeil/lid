<template>
  <div>
    <SearchbarInput />

    <span class="mx-2">
      <small>
        Tracking {{ files.length }} file{{ files.length === 1 ? "" : "s" }} ({{
          totalStorageSpaceUsed
        }})
      </small>
    </span>

    <FileList />
  </div>
</template>

<script>
import { storeToRefs } from "pinia";
import { computed } from "vue";

import SearchbarInput from "@/components/SearchbarInput.vue";
import FileList from "@/components/FileList.vue";
import { useFileStore } from "@/store/fileStore";
import { bytesToHumanReadableUnits } from "../services/dataSizeConversion";

export default {
  name: "Home",
  components: {
    SearchbarInput,
    FileList,
  },
  setup() {
    const fileStore = useFileStore();
    const { files } = storeToRefs(fileStore);
    const totalStorageSpaceUsed = computed(() => {
      return bytesToHumanReadableUnits(
        files.value.reduce((a, c) => a + c.size, 0)
      );
    });

    return {
      files,
      totalStorageSpaceUsed,
    };
  },
};
</script>
