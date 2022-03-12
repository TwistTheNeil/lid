export const bytesToHumanReadableSize = (bytes) => {
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

  return `${gib} GiB`;
};
