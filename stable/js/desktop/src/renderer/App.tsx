import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
import logo from './logo.svg';
import './App.css';

import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import purple from '@material-ui/core/colors/purple';
import MainDrawer from './components/main-drawer';

const theme = createMuiTheme({
  typography: {
    useNextVariants: true
  },
  palette: {
    type: 'light',
    primary: purple,
    secondary: {
      main: '#f44336'
    }
  }
});

const root = {
  display: 'flex'
};

const content = {
  flexGrow: 1,
  backgroundColor: theme.palette.background.default
  // padding: theme.spacing.unit * 3
};

export default function App() {
  return (
    <div style={root}>
      <MuiThemeProvider theme={theme}>
        <MainDrawer />
        <main style={content}>
          <div className="App">
            <header className="App-header">
              <img src={logo} className="App-logo" alt="logo" />
              <p>
                Edit <code>src/App.js</code> and save to reload.
              </p>
              <a
                className="App-link"
                href="https://reactjs.org"
                target="_blank"
                rel="noopener noreferrer"
              >
                Learn React
              </a>
            </header>
          </div>
        </main>
      </MuiThemeProvider>
    </div>
  );
}
