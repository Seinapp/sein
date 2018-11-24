import React, { SFC } from 'react';
import logo from './logo.svg';
import { ActionBar } from './components/action-bar';
import { FilterDrawer } from './components/filter-drawer';
import { withStyles, createStyles, Theme } from '@material-ui/core';

export interface HomeClasses {
  root: any;
  content: any;
}

interface HomeProps {
  classes: HomeClasses;
}

const styles = ({ palette }: Theme) =>
  createStyles({
    content: {
      flexGrow: 1,
      backgroundColor: palette.background.default
    },
    root: {
      display: 'flex'
    }
  });

const HomComponent: SFC<HomeProps> = (props: HomeProps) => {
  const { classes } = props;
  return (
    <>
      <ActionBar />
      <div className={classes.root}>
        <FilterDrawer />
        <main className={classes.content}>
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
      </div>
    </>
  );
};

export const Home = withStyles(styles)(HomComponent);
