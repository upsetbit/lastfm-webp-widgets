const soundBarsAmount = 6;

/* ---------------------------------------------------------------------------------------------- */

const randomInRange = (max, min) => Math.floor(Math.random() * (max - min + 1)) + min;

const setSoundBarHeight = (n, height) => {
  document
    .getElementById('sb' + n)
    .style
    .height = height + '%';
}

const applySequentially = (callback) => {
  for(let i = 1; i <= soundBarsAmount; i++) {
    setSoundBarHeight(i, callback(i));
  }
}

/* ---------------------------------------------------------------------------------------------- */

const hideSoundWaves = () => applySequentially(() => 0);

const randomizeSoundWaves = () => applySequentially((i) => {
  switch(i) {
    case 1: return 40 + randomInRange(0, 35); // 40 - 75
    case 2: return 15 + randomInRange(0, 40); // 15 - 55
    case 3: return 70 + randomInRange(0, 25); // 70 - 95
    case 4: return 65 + randomInRange(0, 20); // 65 - 85
    case 5: return 20 + randomInRange(0, 50); // 20 - 70
    case 6: return 55 + randomInRange(0, 45); // 55 - 100
  }
})
