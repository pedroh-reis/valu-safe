import React from 'react';
import { makeStyles, withStyles } from '@material-ui/core/styles';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import PropTypes from 'prop-types';
import { Colors, MAIN_TEXT_FONT } from '../../../config';
import MapIcon from '@material-ui/icons/Map';
import NotesIcon from '@material-ui/icons/Notes';
import EqualizerIcon from '@material-ui/icons/Equalizer';
import InfoOutlinedIcon from '@material-ui/icons/InfoOutlined';
import LockerOneIcon from '@material-ui/icons/LooksOne';
import LockerTwoIcon from '@material-ui/icons/LooksTwo';

const StyledTabs = withStyles({
  indicator: {
    display: 'flex',
    justifyContent: 'center',
    backgroundColor: 'transparent',
    '& > div': {
      //maxWidth: 40,
      width: '100%',
      backgroundColor: Colors.VALUSAFE_BLUE,
    },
  },
})((props) => <Tabs {...props} TabIndicatorProps={{ children: <div /> }} />);

const StyledTab = withStyles((theme) => ({
  root: {
    textTransform: 'none',
    color: Colors.WHITE,
    fontWeight: theme.typography.fontWeightRegular,
    fontSize: theme.typography.pxToRem(15),
    fontFamily: MAIN_TEXT_FONT,
    '&:hover': {
      color: Colors.VALUSAFE_BLUE,
      opacity: 1,
    },
    '&:focus': {
      opacity: 1,
    },
    '&$selected': {
      color: Colors.VALUSAFE_BLUE,
      fontWeight: theme.typography.fontWeightMedium,
    },
    selected: {},
  },
}))((props) => <Tab disableRipple {...props} />);

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    borderBottom: `1px solid ${Colors.NAVY_BLUE}`,
    backgroundColor: Colors.BLACK,
  },
}));

const TabsBar = (props) => {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <StyledTabs
        variant="fullWidth"
        value={props.value}
        onChange={props.handleTabSelected}
        aria-label="styled tabs example"
      >
        <StyledTab icon={<LockerOneIcon />} />
        <StyledTab icon={<LockerTwoIcon />} />
      </StyledTabs>
    </div>
  );
};

export default TabsBar;

const { func, number } = PropTypes;
TabsBar.propTypes = {
  handleTabSelected: func.isRequired,
  value: number.isRequired,
};
