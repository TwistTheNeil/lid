<template>
  <table class="table table-responsive table-borderless table-hover">
    <thead class="bg-light">
      <tr>
        <th scope="col" width="25%">UUID</th>
        <th scope="col" width="20%">Capacity</th>
        <th scope="col">Name</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="device in filteredStorageDevices" :key="device.uuid">
        <td class="monospace">{{ device.uuid }}</td>
        <td class="monospace" align="right">
          <div class="progress" style="height: 20px">
            <div
              class="progress-bar"
              :class="{
                'bg-danger': device.usagePercentage > 80,
              }"
              role="progressbar"
              :style="{ width: `${device.usagePercentage}%` }"
              :aria-valuenow="device.usagePercentage"
              aria-valuemin="0"
              aria-valuemax="100"
            ></div>
          </div>
        </td>
        <td class="monospace">{{ device.name }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import { storeToRefs } from "pinia";
import { computed } from "vue";
import { Fzf } from "fzf";

import { useStorageDeviceStore } from "@/store/storageDeviceStore";
import { useSearchStore } from "@/store/searchStore";
import { bytesToHumanReadableUnits } from "@/services/dataSizeConversion";

export default {
  name: "DeviceList",
  setup() {
    const storageDeviceStore = useStorageDeviceStore();
    const searchStore = useSearchStore();
    const { devices: storageDevices } = storeToRefs(storageDeviceStore);

    const filteredStorageDevices = computed(() => {
      if (!storageDevices.value || storageDevices.value.length === 0) {
        return [];
      }

      const d = storageDevices.value.map((d) => {
        const capacity = bytesToHumanReadableUnits(d.size);
        const usedCapacityBytes = d.files.reduce((a, c) => a + c.size, 0);
        const usedCapacity = bytesToHumanReadableUnits(usedCapacityBytes);
        const usagePercentage = (usedCapacityBytes / d.size) * 100;

        return {
          ...d,
          capacity,
          usedCapacityBytes,
          usedCapacity,
          usagePercentage,
        };
      });

      if (searchStore.search === "") {
        return d;
      }

      const fzf = new Fzf(d, {
        selector: (i) => `${i.name} ${i.uuid} ${i.capacity}`,
      });

      // We remove whitespace since fzf won't find a whitespace in certain
      // filenames. As a result, weights won't be accurately calculated
      // and expected results will not be seen
      return fzf.find(searchStore.search.replace(/\s/g, "")).map((i) => i.item);
    });

    return {
      filteredStorageDevices,
    };
  },
};
</script>
