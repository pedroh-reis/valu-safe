import { Toolbar, Button, Grid } from '@material-ui/core';
import { StyledAppBar } from './styles';
import PropTypes from 'prop-types';
import React from 'react';
import LogoAppBar from '../../atoms/logoAppBar';
import { Colors } from '../../../config';

const AppBar = (props) => {
  return (
    <div>
      <StyledAppBar position="static">
        <Toolbar>
          <Grid container justify="space-around" alignItems="center">
            <Grid item>
              <LogoAppBar />
            </Grid>
            <Grid item>
              <Button color="inherit" onClick={props.handleLogOutButton}>
                Log out
              </Button>
            </Grid>
          </Grid>
        </Toolbar>
      </StyledAppBar>
      <div style={{ width: '100vw', height: '2px', backgroundColor: Colors.VALUSAFE_BLUE }} />
    </div>
  );
};

export default AppBar;

const {} = PropTypes;
AppBar.propTypes = {};
