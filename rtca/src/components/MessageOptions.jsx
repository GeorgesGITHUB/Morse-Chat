import { Stack, IconButton } from "@mui/joy"
import CachedIcon from '@mui/icons-material/Cached';
import PlayCircleOutlineIcon from '@mui/icons-material/PlayCircleOutline';

function MessageOptions({handleTranslateToggle, handlePlaySound}) {
    return (
        <Stack
            direction="column"
            justifyContent="center"
            alignItems="center"
            spacing={0.01}
        >
            <IconButton 
                size="sm"
                onClick={handleTranslateToggle}
            ><CachedIcon/></IconButton>
            <IconButton 
                size="sm"
                onClick={handlePlaySound}
            ><PlayCircleOutlineIcon/></IconButton>
        </Stack>
    )
}

export default MessageOptions