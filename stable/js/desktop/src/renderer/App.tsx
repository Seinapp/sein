import React from 'react';
import './App.css';

import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import purple from '@material-ui/core/colors/purple';
import MainDrawer from './components/main-drawer';
import { AppRouter } from './components/app-router';

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
};

export default function App() {
  return (
    <div style={root}>
      <MuiThemeProvider theme={theme}>
        <MainDrawer />
        <main style={content}>
          <AppRouter />
        </main>
      </MuiThemeProvider>
    </div>
  );
}
