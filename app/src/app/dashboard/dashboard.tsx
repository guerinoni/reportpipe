'use client';

import * as React from 'react';
import { CssVarsProvider } from '@mui/joy/styles';
import CssBaseline from '@mui/joy/CssBaseline';
import Box from '@mui/joy/Box';
import Button from '@mui/joy/Button';
import Breadcrumbs from '@mui/joy/Breadcrumbs';
import Link from '@mui/joy/Link';
import Typography from '@mui/joy/Typography';

import HomeRoundedIcon from '@mui/icons-material/HomeRounded';
import ChevronRightRoundedIcon from '@mui/icons-material/ChevronRightRounded';
import DownloadRoundedIcon from '@mui/icons-material/DownloadRounded';

import OrderTable from "@/components/ReportTable";
import OrderList from "@/components/OrderList";
import Sidebar from "@/components/Sidebar";
import Header from "@/components/Header";
import {useRouter} from "next/navigation";
import Modal from "@mui/joy/Modal";
import Sheet from "@mui/joy/Sheet";
import ModalClose from "@mui/joy/ModalClose";
import FormControl from "@mui/joy/FormControl";
import FormLabel from "@mui/joy/FormLabel";
import Input from "@mui/joy/Input";
import {useState} from "react";
import Select from '@mui/joy/Select';
import Option from '@mui/joy/Option';

export default function Dashboard() {
    const router = useRouter();
    const [open, setOpen] = React.useState(false);
    const [report, setReport] = useState({})
    const handleReportName = (event: React.SyntheticEvent | null,
                              value: string | null,) => {
        setReport({...report, name: value})
    }

    const handleUserName = (event: React.SyntheticEvent | null,
                            value: Object | null,) => {
        setReport({...report, user: value})
    }
    const handleModalSubmit = () => {
        try {
            //try to create new report then go to dashboard
            //router.push("/dashboard")
        } catch (e) {
            console.log("E: ", e)
        }
    }
  return (
    <CssVarsProvider disableTransitionOnChange>
      <CssBaseline />
      <Box sx={{ display: 'flex', minHeight: '100dvh' }}>
      <Header />
      <Sidebar />
          <Modal
              aria-labelledby="modal-title"
              aria-describedby="modal-desc"
              open={open}
              onClose={() => setOpen(false)}
              sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}
          >
              <Sheet
                  variant="outlined"
                  sx={{
                      maxWidth: 500,
                      borderRadius: 'md',
                      p: 3,
                      boxShadow: 'lg',
                      width: 300,
                      mx: 'auto',
                      my: 4,
                      py: 3,
                      px: 2,
                      display: 'flex',
                      flexDirection: 'column',
                      gap: 2,
                  }}
              >
                  <ModalClose variant="plain" sx={{ m: 1 }} />
                  <Typography
                      component="h2"
                      id="modal-title"
                      level="h4"
                      textColor="inherit"
                      fontWeight="lg"
                      mb={1}
                  >
                      Create report
                  </Typography>
                  <FormControl id="reportName">
                      <FormLabel>Report Name</FormLabel>
                      <Input name="reportName" type="text" placeholder="New Report" onChange={handleReportName}/>
                  </FormControl>
                  <FormControl id="userId">
                      <FormLabel>User account</FormLabel>
                      <Select defaultValue="dog" onChange={handleUserName}>
                          <Option value="dog">Dog</Option>
                          <Option value="cat">Cat</Option>
                          <Option value="fish">Fish</Option>
                          <Option value="bird">Bird</Option>
                      </Select>
                  </FormControl>
                  <Button sx={{ mt: 1 }} onClick={handleModalSubmit}>Create report</Button>
              </Sheet>
          </Modal>
        <Box
          component="main"
          className="MainContent"
          sx={{
            px: { xs: 2, md: 6 },
            pt: {
              xs: 'calc(12px + var(--Header-height))',
              sm: 'calc(12px + var(--Header-height))',
              md: 3,
            },
            pb: { xs: 2, sm: 2, md: 3 },
            flex: 1,
            display: 'flex',
            flexDirection: 'column',
            minWidth: 0,
            height: '100dvh',
            gap: 1,
          }}
        >
          <Box sx={{ display: 'flex', alignItems: 'center' }}>
            <Breadcrumbs
              size="sm"
              aria-label="breadcrumbs"
              separator={<ChevronRightRoundedIcon />}
              sx={{ pl: 0 }}
            >
              <Link
                underline="none"
                color="neutral"
                href="#some-link"
                aria-label="Home"
              >
                <HomeRoundedIcon />
              </Link>
              <Link
                underline="hover"
                color="neutral"
                href="#some-link"
                fontSize={12}
                fontWeight={500}
              >
                Dashboard
              </Link>
              <Typography color="primary" fontWeight={500} fontSize={12}>
                Reports
              </Typography>
            </Breadcrumbs>
          </Box>
          <Box
            sx={{
              display: 'flex',
              mb: 1,
              gap: 1,
              flexDirection: { xs: 'column', sm: 'row' },
              alignItems: { xs: 'start', sm: 'center' },
              flexWrap: 'wrap',
              justifyContent: 'space-between',
            }}
          >
            <Typography level="h2" component="h1">
              Reports
            </Typography>
            <Button
              color="primary"
              size="sm"
              onClick={() => {
                  setOpen(true)}}
            >
              Crea Report
            </Button>
          </Box>
          <OrderTable />
          <OrderList />
        </Box>
      </Box>
    </CssVarsProvider>
  );
}
