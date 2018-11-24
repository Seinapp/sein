import React, { SFC } from 'react';
import { Theme, withStyles, createStyles } from '@material-ui/core/styles';

import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import InputBase from '@material-ui/core/InputBase';

import AddIcon from '@material-ui/icons/Add';
import SettingsIcon from '@material-ui/icons/Settings';
import SyncIcon from '@material-ui/icons/Sync';
import SyncIssueIcon from '@material-ui/icons/SyncProblem';
import SearchIcon from '@material-ui/icons/Search';
import { fade } from '@material-ui/core/styles/colorManipulator';

interface ActionBarClasses {
  root: any;
  appBar: any;
  toolbar: any;
  grow: any;
  search: any;
  searchIcon: any;
  inputRoot: any;
  inputInput: any;
}

interface ActionBarProps {
  classes: ActionBarClasses;
}

const styles = ({
  palette,
  spacing,
  shape,
  breakpoints,
  transitions,
  zIndex
}: Theme) =>
  createStyles({
    appBar: {
      WebkitAppRegion: 'drag',
      zIndex: zIndex.drawer + 1
    },
    toolbar: {
      paddingLeft: '64px'
    },
    root: {
      flexGrow: 1
    },
    grow: {
      flexGrow: 1
    },
    search: {
      position: 'relative',
      borderRadius: shape.borderRadius,
      backgroundColor: fade(palette.common.white, 0.15),
      '&:hover': {
        backgroundColor: fade(palette.common.white, 0.25)
      },
      marginRight: spacing.unit * 2,
      marginLeft: 0,
      width: '100%',
      [breakpoints.up('sm')]: {
        marginLeft: spacing.unit * 3,
        width: 'auto'
      }
    },
    searchIcon: {
      width: spacing.unit * 9,
      height: '100%',
      position: 'absolute',
      pointerEvents: 'none',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center'
    },
    inputRoot: {
      color: 'inherit',
      width: '100%'
    },
    inputInput: {
      paddingTop: spacing.unit,
      paddingRight: spacing.unit,
      paddingBottom: spacing.unit,
      paddingLeft: spacing.unit * 10,
      transition: transitions.create('width'),
      width: '100%',
      [breakpoints.up('md')]: {
        width: 200
      }
    }
  });

const ActionBarCmp: SFC<ActionBarProps> = (props: ActionBarProps) => {
  const { classes } = props;

  return (
    <div className={classes.root}>
      <AppBar className={classes.appBar} position="fixed">
        <Toolbar variant="dense" className={classes.toolbar}>
          <div className={classes.search}>
            <div className={classes.searchIcon}>
              <SearchIcon />
            </div>
            <InputBase
              placeholder="Search…"
              classes={{
                root: classes.inputRoot,
                input: classes.inputInput
              }}
            />
          </div>

          <IconButton color="inherit">
            <AddIcon />
          </IconButton>

          <div className={classes.grow} />

          <IconButton color="inherit">
            <SyncIcon />
          </IconButton>

          <IconButton color="inherit">
            <SyncIssueIcon />
          </IconButton>

          <IconButton color="inherit">
            <SettingsIcon />
          </IconButton>
        </Toolbar>
      </AppBar>
    </div>
  );
};

export const ActionBar = withStyles(styles)(ActionBarCmp);
