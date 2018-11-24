import React from 'react';
import './App.css';

import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import purple from '@material-ui/core/colors/purple';
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

export default function App() {
  return (
    <MuiThemeProvider theme={theme}>
      <AppRouter />
    </MuiThemeProvider>
  );
}
