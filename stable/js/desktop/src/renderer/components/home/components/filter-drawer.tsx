import React, { SFC } from 'react';

import Drawer from '@material-ui/core/Drawer';
import { Theme, withStyles, createStyles } from '@material-ui/core/styles';
import { fade } from '@material-ui/core/styles/colorManipulator';

export interface FilterDrawerClasses {
  drawer: any;
  drawerPaper: any;
}

export interface FilterDrawerProps {
  classes: FilterDrawerClasses;
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

const FilterDrawerCmp: SFC<FilterDrawerProps> = (props: FilterDrawerProps) => {
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
};

export const FilterDrawer = withStyles(styles)(FilterDrawerCmp);
