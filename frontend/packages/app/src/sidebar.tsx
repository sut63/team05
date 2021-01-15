import React, { useState, useEffect } from 'react';
import HomeIcon from '@material-ui/icons/Home';
import SignOut from '@material-ui/icons/Settings';
import MonetizationOnIcon from '@material-ui/icons/MonetizationOn';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';
import PermIdentityIcon from '@material-ui/icons/PermIdentity';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import MeetingRoomIcon from '@material-ui/icons/MeetingRoom';
import CommentIcon from '@material-ui/icons/Comment';

import PlaylistAddIcon from '@material-ui/icons/PlaylistAdd';


import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarPinButton,
  SidebarThemeToggle,
} from '@backstage/core';

import { EntMember } from 'plugin-welcome/src/api/models/EntMember';
import { DefaultApi } from 'plugin-welcome/src/api/apis';
import { EntOfficer } from 'plugin-welcome/src/api';
export const AppSidebar = () => {

  const api = new DefaultApi();
  const [memberid, setMember] = useState(Number);
  const [members, setMembers] = useState<EntMember[]>([]);
  const [loading, setLoading] = useState(true);

  const [officers, setOfficers] = useState<EntOfficer[]>([]);
  const [officerID, setOfficerID] = useState(Number);

  useEffect(() => {
    const getMembers = async () => {
      const res = await api.listMember({});
      setLoading(false);
      setMembers(res);
    };
    getMembers();
    const data = localStorage.getItem("memberdata");
    if (data) {
      setMember(Number(JSON.parse(data)));
      setLoading(false);
    }

    const getOfficers = async () => {
      const ofc: any = await api.listOfficer({});
      setLoading(false);
      setOfficers(ofc);
    }
    getOfficers();
    const dataa = localStorage.getItem("officerdata");
    if (dataa) {
      setOfficerID(Number(JSON.parse(dataa)));
      setLoading(false);
    }
    
  }, [loading]);

  return (

    <Sidebar>
      <SidebarDivider />
      {/* Global nav, not org-specific */}
      {(memberid) ?
        members.filter((filter:EntMember) => filter.id == memberid).map((item:EntMember) => 
          <><SidebarItem icon={PermIdentityIcon} text={item.memberName} />
            <SidebarItem
            icon={ShoppingCartIcon}
            to="Insurance"
            text="Insurance" />
            <SidebarItem
            icon={MonetizationOnIcon}
            to="payment"
            text="Payment" />
            <SidebarItem
            icon={CommentIcon}
            to="Inquiry"
            text="Inquiry" />
            </>
        )
        :
        null
      }

      {/* End global nav */}

      {/* Global nav, not org-specific */}
      {(officerID) ?
        officers.filter((filter:EntOfficer) => filter.id == officerID).map((item:EntOfficer) => 
          <><SidebarItem icon={PermIdentityIcon} text={item.officerName} />
          <SidebarItem
            icon={PlaylistAddIcon}
            to="Product"
            text="Product" />
            </>
        )
        :
        null
      }

      {/* End global nav */}

      
      
      <SidebarDivider />
      <SidebarSpace />
      <SidebarDivider />
      <SidebarThemeToggle />
      {(memberid) ?
        <SidebarItem icon={MeetingRoomIcon} to="./SignIn" text="ออกจากระบบ"
          onClick={() => {
            localStorage.setItem("memberdata", JSON.stringify(null));
            history.pushState("", "", "./SignIn");
            window.location.reload(false);
          }} />
        :
        null
      }

      {(officerID) ?
        <SidebarItem icon={MeetingRoomIcon} to="./Officerlogin" text="ออกจากระบบ"
          onClick={() => {
            localStorage.setItem("officerdata", JSON.stringify(null));
            history.pushState("", "", "./Officerlogin");
            window.location.reload(false);
          }} />
        :
        null
      }
      <SidebarPinButton />
    </Sidebar>
  )
};