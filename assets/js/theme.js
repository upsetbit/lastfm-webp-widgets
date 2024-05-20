class WidgetTheme {
  constructor() {
    this.root = document.documentElement;
    this.isDarkMode = false;
  }

  toggle() {
    this.isDarkMode = !this.isDarkMode;
    this.root.setAttribute('class', this.isDarkMode ? 'dark-mode' : '')
  }

  set(mode) {
    if (mode === 'dark') {
      this.isDarkMode = true;
      this.root.setAttribute('class', 'dark-mode');
      return
    }

    if (mode === 'light') {
      this.isDarkMode = false;
      this.root.setAttribute('class', '');
      return
    }

    throw new Error('unknown theme mode');
  }
}

const theme = new WidgetTheme();
