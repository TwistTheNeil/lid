export const bytesToHumanReadableUnits = (bytes) => {
  const units = ['B', 'KiB', 'MiB', 'GiB', 'TiB'];
  const base = 1024;

  const logArg = Math.floor(Math.log(bytes)/Math.log(base));
  const unit = units[logArg];
  const fixedValue = (Math.floor(bytes/Math.pow(base, logArg))).toFixed(2);


  return `${fixedValue} ${unit}`;
};
