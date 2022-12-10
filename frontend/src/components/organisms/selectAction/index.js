import React from 'react';
import SelectScreenButton from '../../atoms/selectScreenButton';
import SettingsRemoteIcon from '@material-ui/icons/SettingsRemote';
import MonitorIcon from '@material-ui/icons/Timeline';
import { Colors, MAIN_TEXT_FONT } from '../../../config';
import PropTypes from 'prop-types';

const SelectAction = (props) => {
  return (
    <div
      style={{
        color: Colors.WHITE,
        flexDirection: 'column',
        display: 'flex',
        justifyContent: 'space-evenly',
        alignItems: 'center',
        height: '80vh',
      }}
    >
      <SelectScreenButton onClick={props.handleControlSelected}>
        <SettingsRemoteIcon style={{ color: Colors.VALUSAFE_BLUE, marginRight: '5px' }} />
        <div style={{ fontFamily: MAIN_TEXT_FONT, color: Colors.VALUSAFE_BLUE }}>Controlar</div>
      </SelectScreenButton>
      <div style={{ height: '2px', backgroundColor: Colors.NAVY_BLUE, width: '-webkit-fill-available' }} />
      <SelectScreenButton onClick={props.handleMonitorSelected}>
        <MonitorIcon style={{ color: Colors.VALUSAFE_BLUE, marginRight: '5px' }} />
        <div style={{ fontFamily: MAIN_TEXT_FONT, color: Colors.VALUSAFE_BLUE }}>Monitorar</div>
      </SelectScreenButton>
    </div>
  );
};

export default SelectAction;

const { func } = PropTypes;
SelectAction.propTypes = {
  handleControlSelected: func.isRequired,
  handleMonitorSelected: func.isRequired,
};
