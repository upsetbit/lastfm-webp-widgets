const track = document.getElementById('track-title');
const scrollJump = 40;

/* ---------------------------------------------------------------------------------------------- */

const measureTrackTitleInPixels = () => {
  const canvas = document.createElement('canvas');

  const ctx = canvas.getContext('2d');
  ctx.font = window.getComputedStyle(track).font;

  return ctx.measureText(track.textContent).width;
}

const calcRequiredScrollTicks = () => Math.ceil((Math.ceil(measureTrackTitleInPixels() / scrollJump) + 5) / 2);

/* ---------------------------------------------------------------------------------------------- */

const requiredScrollTicks = calcRequiredScrollTicks();

let tickCounter = requiredScrollTicks;

const tickScroll = () => {
  if (tickCounter === 0) {
    track.scroll(0, 0);

    tickCounter = requiredScrollTicks;
    return
  }

  tickCounter--;
  track.scroll((requiredScrollTicks - tickCounter) * scrollJump, 0);
}
