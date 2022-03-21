<template>
  <div>
    <SearchbarInput />

    <div v-if="queriedStorageDevice">
      <span class="mx-2">
        <small>
          Tracking {{ queriedStorageDevice.files.length }} file{{
            queriedStorageDevice.files.length === 1 ? "" : "s"
          }}
          (
          {{ totalStorageSpaceUsed }}
          )
        </small>
      </span>

      <FileList :files="queriedStorageDevice.files" />
    </div>
  </div>
</template>

<script>
import { storeToRefs } from "pinia";
import { useRoute } from "vue-router";
import { computed } from "vue";

import SearchbarInput from "@/components/SearchbarInput.vue";
import { useStorageDeviceStore } from "@/store/storageDeviceStore";
import FileList from "@/components/FileList.vue";
import { bytesToHumanReadableUnits } from "@/services/dataSizeConversion";

export default {
  name: "DeviceFileList",
  components: {
    SearchbarInput,
    FileList,
  },
  setup() {
    const route = useRoute();
    const storageDeviceStore = useStorageDeviceStore();

    const { devices: storageDevices } = storeToRefs(storageDeviceStore);
    const queriedStorageDevice = computed(() =>
      storageDevices.value.find((d) => d.uuid === route.params.uuid)
    );

    const totalStorageSpaceUsed = computed(() =>
      bytesToHumanReadableUnits(
        queriedStorageDevice.value.files.reduce((a, c) => a + c.size, 0)
      )
    );

    return {
      storageDevices,
      queriedStorageDevice,
      totalStorageSpaceUsed,
    };
  },
};
</script>
