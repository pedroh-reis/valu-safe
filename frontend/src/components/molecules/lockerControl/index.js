import React from 'react';
import SelectScreenButton from '../../atoms/selectScreenButton';
import SettingsRemoteIcon from '@material-ui/icons/SettingsRemote';
import MonitorIcon from '@material-ui/icons/Timeline';
import { Colors, MAIN_TEXT_FONT } from '../../../config';
import PropTypes from 'prop-types';
import { Button } from '@mui/material';

const LockerControl = (props) => {
  return (
    <div
      style={{
        color: Colors.WHITE,
        flexDirection: 'column',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '74vh',
      }}
    >
      <h2>{`Controle - Locker ${props.lockerId + 1}`}</h2>
      <h4 style={{ marginBottom: '54px', color: props.lockerUnlocked ? Colors.GREEN : Colors.RED }}>{`${
        props.lockerUnlocked ? 'Aberto' : 'Fechado'
      }`}</h4>
      <Button
        variant="contained"
        onClick={props.handleChangeLockerStatus}
        style={{ backgroundColor: props.lockerUnlocked ? Colors.RED : Colors.GREEN, color: Colors.BLACK }}
      >
        <strong>{`${props.lockerUnlocked ? 'Fechar' : 'Abrir'}`}</strong>
      </Button>
    </div>
  );
};

export default LockerControl;

const { func, number, bool } = PropTypes;
LockerControl.propTypes = {
  lockerId: number.isRequired,
  handleChangeLockerStatus: func.isRequired,
  lockerUnlocked: bool.isRequired,
};
