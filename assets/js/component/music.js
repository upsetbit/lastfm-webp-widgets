class Music {
  constructor() {
    this.scrollJump = 40;

    this.trackTitle = document.getElementById('track-title');
    this.artistName = document.getElementById('artist-name');
    this.albumCover = document.getElementById('album-cover');

    this.refreshMeasurements();
  }

  refreshMeasurements() {
    this.trackTitleSizeInPixels = this.measureTrackTitleInPixels();
    this.requiredScrollTicks    = this.calcRequiredScrollTicks();
    this.tickCounter            = this.requiredScrollTicks;

    return this.trackTitleSizeInPixels;
  }

  measureTrackTitleInPixels() {
    const canvas = document.createElement('canvas');

    const ctx = canvas.getContext('2d');
    ctx.font = window.getComputedStyle(this.trackTitle).font;

    return ctx.measureText(this.trackTitle.textContent).width;
  }

  calcRequiredScrollTicks() {
    return Math.ceil((Math.ceil(this.trackTitleSizeInPixels / this.scrollJump) + 5) / 2);
  }

  tickScroll() {
    if (this.tickCounter === 0) {
      this.trackTitle.scroll(0, 0);

      this.tickCounter = this.requiredScrollTicks;
      return
    }

    this.tickCounter--;
    this.trackTitle.scroll((this.requiredScrollTicks - this.tickCounter) * this.scrollJump, 0);

    return this.tickCounter;
  }

  /* -------------------------------------------------------------------------------------------- */

  setTrackTitle(title) {
    this.trackTitle.textContent = title;
    this.trackTitle.scroll(0, 0);
    this.refreshMeasurements();

    return title;
  }

  getTrackTitleSizeInPixels() {
    return this.trackTitleSizeInPixels;
  }

  getTrackTitleSizeInPixelsRounded() {
    return Math.ceil(this.trackTitleSizeInPixels);
  }

  getScrollTicksAmount() {
    return this.requiredScrollTicks;
  }

  /* -------------------------------------------------------------------------------------------- */

  setArtistName(name) {
    this.artistName.textContent = name;
    return name;
  }

  setAlbumCoverSource(src) {
    this.albumCover.src = src;
    return src;
  }
}

const music = new Music();
