import React, { useState, useEffect } from "react";
import { DragDropContext } from "react-beautiful-dnd";
import Column from "./Column";
import { resetServerContext } from "react-beautiful-dnd"
import CssBaseline from "@mui/joy/CssBaseline";
import Header from "@/components/Header";
import Sidebar from "@/components/Sidebar";
import Box from "@mui/joy/Box";
import {CssVarsProvider} from "@mui/joy/styles";

export default function KanbanBoard() {
    const [completed, setCompleted] = useState([]);
    const [incomplete, setIncomplete] = useState([]);
    const [backlog, setBacklog] = useState([]);
    const [inReview, setInReview] = useState([]);

    resetServerContext()


    useEffect(() => {
        fetch("https://jsonplaceholder.typicode.com/todos")
            .then((response) => response.json())
            .then((json) => {
                setCompleted(json.filter((task) => task.completed));
                setIncomplete(json.filter((task) => !task.completed));
            });
    }, []);

    const handleDragEnd = (result) => {
        const { destination, source, draggableId } = result;

        if (!destination || source.droppableId === destination.droppableId) return;

        deletePreviousState(source.droppableId, draggableId);

        const task = findItemById(draggableId, [...incomplete, ...completed, ...inReview, ...backlog]);

        setNewState(destination.droppableId, task);

    };

    function deletePreviousState(sourceDroppableId, taskId) {
        switch (sourceDroppableId) {
            case "1":
                setIncomplete(removeItemById(taskId, incomplete));
                break;
            case "2":
                setCompleted(removeItemById(taskId, completed));
                break;
            case "3":
                setInReview(removeItemById(taskId, inReview));
                break;
            case "4":
                setBacklog(removeItemById(taskId, backlog));
                break;
        }

    }
    function setNewState(destinationDroppableId, task) {
        let updatedTask;
        switch (destinationDroppableId) {
            case "1":   // TO DO
                updatedTask = { ...task, completed: false };
                setIncomplete([updatedTask, ...incomplete]);
                break;
            case "2":  // DONE
                updatedTask = { ...task, completed: true };
                setCompleted([updatedTask, ...completed]);
                break;
            case "3":  // IN REVIEW
                updatedTask = { ...task, completed: false };
                setInReview([updatedTask, ...inReview]);
                break;
            case "4":  // BACKLOG
                updatedTask = { ...task, completed: false };
                setBacklog([updatedTask, ...backlog]);
                break;
        }
    }
    function findItemById(id, array) {
        return array.find((item) => item.id == id);
    }

    function removeItemById(id, array) {
        return array.filter((item) => item.id != id);
    }

    return (
        <CssVarsProvider disableTransitionOnChange>
            <CssBaseline />
            <Box sx={{ display: 'flex', minHeight: '100dvh' }}>
                <Header />
                <Sidebar />
                <DragDropContext onDragEnd={handleDragEnd}>
                    <Box
                        sx={{
                            display: "flex",
                        }}
                    >
                        <Column title={"TO DO"} tasks={incomplete} id={"1"} />
                        <Column title={"DONE"} tasks={completed} id={"2"} />
                        <Column title={"IN REVIEW"} tasks={inReview} id={"3"} />
                        <Column title={"BACKLOG"} tasks={backlog} id={"4"} />
                    </Box>
                </DragDropContext>
            </Box>
        </CssVarsProvider>
    );
}