import React, { ReactNode, SFC } from 'react';
import { HashRouter as Router, Route, Link } from 'react-router-dom';
import { Home } from './home';

interface AppRouterProps {
  children?: ReactNode;
}

// Using HashRouter here (index.html#!/) because otherwise the current
// path (window.location) is something like:
// file:///path/to/the/app/index.html

export const AppRouter: SFC<AppRouterProps> = (props: AppRouterProps) => {
  return (
    <Router>
      <>
        {props.children}

        <Route exact path="/" component={Home} />
      </>
    </Router>
  );
};
