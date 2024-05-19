class SoundWaves {
  constructor() {
    this.soundBarsAmount = 6;
  }

  randomInRange(max, min) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }

  setSoundBarHeight(n, height) {
    document
      .getElementById('sb' + n)
      .style
      .height = height + '%';
  }

  applySequentially(callback) {
    for(let i = 1; i <= this.soundBarsAmount; i++) {
      this.setSoundBarHeight(i, callback(i));
    }
  }

  hide() {
    this.applySequentially(() => 0);
  }

  randomize() {
    this.applySequentially((i) => {
      switch(i) {
        case 1: return 40 + this.randomInRange(0, 35); // 40 - 75
        case 2: return 15 + this.randomInRange(0, 40); // 15 - 55
        case 3: return 70 + this.randomInRange(0, 25); // 70 - 95
        case 4: return 65 + this.randomInRange(0, 20); // 65 - 85
        case 5: return 20 + this.randomInRange(0, 50); // 20 - 70
        case 6: return 55 + this.randomInRange(0, 45); // 55 - 100
      }
    });
  }
}

const waves = new SoundWaves();
