<template>
  <div>
    <SearchbarInput />

    <span class="mx-2">
      <small>
        Tracking {{ storageDevices.length }} storage device{{
          storageDevices.length === 1 ? "" : "s"
        }}
      </small>
    </span>

    <DeviceList />
  </div>
</template>

<script>
import { storeToRefs } from "pinia";

import SearchbarInput from "@/components/SearchbarInput.vue";
import DeviceList from "@/components/DeviceList.vue";
import { useStorageDeviceStore } from "@/store/storageDeviceStore";

export default {
  name: "Devices",
  components: {
    SearchbarInput,
    DeviceList,
  },
  setup() {
    const storageDeviceStore = useStorageDeviceStore();
    const { devices: storageDevices } = storeToRefs(storageDeviceStore);
    return {
      storageDevices,
    };
  },
};
</script>
