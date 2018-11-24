import React, { ReactElement } from 'react';

import Drawer from '@material-ui/core/Drawer';
import { Theme, withStyles, createStyles } from '@material-ui/core/styles';
import { fade } from '@material-ui/core/styles/colorManipulator';

export interface MainDrawerClasses {
  drawer: any;
  drawerPaper: any;
}

export interface MainDrawerProps {
  classes: MainDrawerClasses;
}

const drawerWidth = 240;
const styles = ({ palette, spacing }: Theme) =>
  createStyles({
    drawer: {
      width: drawerWidth,
      flexShrink: 0
    },
    drawerPaper: {
      width: drawerWidth,
      backgroundColor: fade(palette.background.paper, 0.5)
    }
  });

function MainDrawer(props: MainDrawerProps): ReactElement<any> {
  const { classes } = props;

  return (
    <Drawer
      className={classes.drawer}
      variant="permanent"
      classes={{
        paper: classes.drawerPaper
      }}
      anchor="left"
    >
      {/*  */}
    </Drawer>
  );
}

export default withStyles(styles)(MainDrawer);
