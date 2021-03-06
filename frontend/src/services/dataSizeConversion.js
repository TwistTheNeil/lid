export const bytesToHumanReadableUnits = (bytes) => {
  const kib = (bytes / 1024).toFixed(2);
  if (kib < 1) {
    return `${bytes} B`;
  }

  const mib = (kib / 1024).toFixed(2);
  if (mib < 1) {
    return `${kib} KiB`;
  }

  const gib = (mib / 1024).toFixed(2);
  if (gib < 1) {
    return `${mib} MiB`;
  }

  const tib = (gib / 1024).toFixed(2);
  if (tib < 1) {
    return `${gib} GiB`;
  }

  return `${tib} TiB`;
};
